package container

import (
	"context"
	"fmt"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
	v1 "k8s.io/api/core/v1"
)

const k8sNodeName = "kubernetes_nodes"

func gatherNode(client k8sClientX, extraTags map[string]string) (*k8sResourceStats, error) {
	list, err := client.getNodes().List(context.Background(), metaV1ListOption)
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes resource: %w", err)
	}

	if len(list.Items) == 0 {
		return nil, nil
	}
	return exportNode(list.Items, extraTags), nil
}

func exportNode(items []v1.Node, extraTags tagsType) *k8sResourceStats {
	res := newK8sResourceStats()

	for _, item := range items {
		obj := newNode()
		obj.tags["name"] = fmt.Sprintf("%v", item.UID)
		obj.tags["node_name"] = item.Name
		obj.tags["status"] = fmt.Sprintf("%v", item.Status.Phase)

		if _, ok := item.Labels["node-role.kubernetes.io/master"]; ok {
			obj.tags["role"] = "master"
		} else {
			obj.tags["role"] = "node"
		}

		obj.tags.addValueIfNotEmpty("cluster_name", defaultClusterName(item.ClusterName))
		obj.tags.addValueIfNotEmpty("namespace", defaultNamespace(item.Namespace))
		for _, address := range item.Status.Addresses {
			if address.Type == v1.NodeInternalIP {
				obj.tags.addValueIfNotEmpty("internal_ip", address.Address)
				obj.tags.addValueIfNotEmpty("node_ip", address.Address) // depercated
			}
		}
		obj.tags.append(extraTags)

		obj.fields["age"] = int64(time.Since(item.CreationTimestamp.Time).Seconds())
		obj.fields["kubelet_version"] = item.Status.NodeInfo.KubeletVersion

		obj.fields.addMapWithJSON("annotations", item.Annotations)
		obj.fields.addLabel(item.Labels)
		obj.fields.mergeToMessage(obj.tags)
		delete(obj.fields, "annotations")

		obj.time = time.Now()
		res.meas = append(res.meas, obj)
	}
	return res
}

type node struct {
	tags   tagsType
	fields fieldsType
	time   time.Time
}

func newNode() *node {
	return &node{
		tags:   make(tagsType),
		fields: make(fieldsType),
	}
}

func (n *node) LineProto() (*io.Point, error) {
	return io.NewPoint(k8sNodeName, n.tags, n.fields, &io.PointOption{Time: n.time, Category: datakit.Object})
}

//nolint:lll
func (*node) Info() *inputs.MeasurementInfo {
	return &inputs.MeasurementInfo{
		Name: k8sNodeName,
		Desc: "Kubernetes node 对象数据",
		Type: "object",
		Tags: map[string]interface{}{
			"name":         inputs.NewTagInfo("UID"),
			"node_name":    inputs.NewTagInfo("Name must be unique within a namespace."),
			"node_ip":      inputs.NewTagInfo("Node IP (depercated)"),
			"internal_ip":  inputs.NewTagInfo("Node internal IP"),
			"role":         inputs.NewTagInfo("Node role. (master/node)"),
			"cluster_name": inputs.NewTagInfo("The name of the cluster which the object belongs to."),
			"namespace":    inputs.NewTagInfo("Namespace defines the space within each name must be unique."),
			"status":       inputs.NewTagInfo("NodePhase is the recently observed lifecycle phase of the node. (Pending/Running/Terminated)"),
		},
		Fields: map[string]interface{}{
			"age":             &inputs.FieldInfo{DataType: inputs.Int, Unit: inputs.DurationSecond, Desc: "age (seconds)"},
			"kubelet_version": &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: "Kubelet Version reported by the node."},
			"message":         &inputs.FieldInfo{DataType: inputs.String, Unit: inputs.UnknownUnit, Desc: "object details"},
		},
	}
}

//nolint:gochecknoinits
func init() {
	registerMeasurement(&node{})
}
