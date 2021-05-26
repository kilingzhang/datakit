// Package storage provides the Chrome DevTools Protocol
// commands, types, and events for the Storage domain.
//
// Generated by the cdproto-gen command.
package storage

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
)

// ClearDataForOriginParams clears storage for origin.
type ClearDataForOriginParams struct {
	Origin       string `json:"origin"`       // Security origin.
	StorageTypes string `json:"storageTypes"` // Comma separated list of StorageType to clear.
}

// ClearDataForOrigin clears storage for origin.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearDataForOrigin
//
// parameters:
//   origin - Security origin.
//   storageTypes - Comma separated list of StorageType to clear.
func ClearDataForOrigin(origin string, storageTypes string) *ClearDataForOriginParams {
	return &ClearDataForOriginParams{
		Origin:       origin,
		StorageTypes: storageTypes,
	}
}

// Do executes Storage.clearDataForOrigin against the provided context.
func (p *ClearDataForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearDataForOrigin, p, nil)
}

// GetCookiesParams returns all browser cookies.
type GetCookiesParams struct {
	BrowserContextID cdp.BrowserContextID `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// GetCookies returns all browser cookies.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getCookies
//
// parameters:
func GetCookies() *GetCookiesParams {
	return &GetCookiesParams{}
}

// WithBrowserContextID browser context to use when called on the browser
// endpoint.
func (p GetCookiesParams) WithBrowserContextID(browserContextID cdp.BrowserContextID) *GetCookiesParams {
	p.BrowserContextID = browserContextID
	return &p
}

// GetCookiesReturns return values.
type GetCookiesReturns struct {
	Cookies []*network.Cookie `json:"cookies,omitempty"` // Array of cookie objects.
}

// Do executes Storage.getCookies against the provided context.
//
// returns:
//   cookies - Array of cookie objects.
func (p *GetCookiesParams) Do(ctx context.Context) (cookies []*network.Cookie, err error) {
	// execute
	var res GetCookiesReturns
	err = cdp.Execute(ctx, CommandGetCookies, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Cookies, nil
}

// SetCookiesParams sets given cookies.
type SetCookiesParams struct {
	Cookies          []*network.CookieParam `json:"cookies"`                    // Cookies to be set.
	BrowserContextID cdp.BrowserContextID   `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// SetCookies sets given cookies.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-setCookies
//
// parameters:
//   cookies - Cookies to be set.
func SetCookies(cookies []*network.CookieParam) *SetCookiesParams {
	return &SetCookiesParams{
		Cookies: cookies,
	}
}

// WithBrowserContextID browser context to use when called on the browser
// endpoint.
func (p SetCookiesParams) WithBrowserContextID(browserContextID cdp.BrowserContextID) *SetCookiesParams {
	p.BrowserContextID = browserContextID
	return &p
}

// Do executes Storage.setCookies against the provided context.
func (p *SetCookiesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetCookies, p, nil)
}

// ClearCookiesParams clears cookies.
type ClearCookiesParams struct {
	BrowserContextID cdp.BrowserContextID `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// ClearCookies clears cookies.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearCookies
//
// parameters:
func ClearCookies() *ClearCookiesParams {
	return &ClearCookiesParams{}
}

// WithBrowserContextID browser context to use when called on the browser
// endpoint.
func (p ClearCookiesParams) WithBrowserContextID(browserContextID cdp.BrowserContextID) *ClearCookiesParams {
	p.BrowserContextID = browserContextID
	return &p
}

// Do executes Storage.clearCookies against the provided context.
func (p *ClearCookiesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearCookies, p, nil)
}

// GetUsageAndQuotaParams returns usage and quota in bytes.
type GetUsageAndQuotaParams struct {
	Origin string `json:"origin"` // Security origin.
}

// GetUsageAndQuota returns usage and quota in bytes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getUsageAndQuota
//
// parameters:
//   origin - Security origin.
func GetUsageAndQuota(origin string) *GetUsageAndQuotaParams {
	return &GetUsageAndQuotaParams{
		Origin: origin,
	}
}

// GetUsageAndQuotaReturns return values.
type GetUsageAndQuotaReturns struct {
	Usage          float64         `json:"usage,omitempty"`          // Storage usage (bytes).
	Quota          float64         `json:"quota,omitempty"`          // Storage quota (bytes).
	OverrideActive bool            `json:"overrideActive,omitempty"` // Whether or not the origin has an active storage quota override
	UsageBreakdown []*UsageForType `json:"usageBreakdown,omitempty"` // Storage usage per type (bytes).
}

// Do executes Storage.getUsageAndQuota against the provided context.
//
// returns:
//   usage - Storage usage (bytes).
//   quota - Storage quota (bytes).
//   overrideActive - Whether or not the origin has an active storage quota override
//   usageBreakdown - Storage usage per type (bytes).
func (p *GetUsageAndQuotaParams) Do(ctx context.Context) (usage float64, quota float64, overrideActive bool, usageBreakdown []*UsageForType, err error) {
	// execute
	var res GetUsageAndQuotaReturns
	err = cdp.Execute(ctx, CommandGetUsageAndQuota, p, &res)
	if err != nil {
		return 0, 0, false, nil, err
	}

	return res.Usage, res.Quota, res.OverrideActive, res.UsageBreakdown, nil
}

// OverrideQuotaForOriginParams override quota for the specified origin.
type OverrideQuotaForOriginParams struct {
	Origin    string  `json:"origin"`              // Security origin.
	QuotaSize float64 `json:"quotaSize,omitempty"` // The quota size (in bytes) to override the original quota with. If this is called multiple times, the overridden quota will be equal to the quotaSize provided in the final call. If this is called without specifying a quotaSize, the quota will be reset to the default value for the specified origin. If this is called multiple times with different origins, the override will be maintained for each origin until it is disabled (called without a quotaSize).
}

// OverrideQuotaForOrigin override quota for the specified origin.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-overrideQuotaForOrigin
//
// parameters:
//   origin - Security origin.
func OverrideQuotaForOrigin(origin string) *OverrideQuotaForOriginParams {
	return &OverrideQuotaForOriginParams{
		Origin: origin,
	}
}

// WithQuotaSize the quota size (in bytes) to override the original quota
// with. If this is called multiple times, the overridden quota will be equal to
// the quotaSize provided in the final call. If this is called without
// specifying a quotaSize, the quota will be reset to the default value for the
// specified origin. If this is called multiple times with different origins,
// the override will be maintained for each origin until it is disabled (called
// without a quotaSize).
func (p OverrideQuotaForOriginParams) WithQuotaSize(quotaSize float64) *OverrideQuotaForOriginParams {
	p.QuotaSize = quotaSize
	return &p
}

// Do executes Storage.overrideQuotaForOrigin against the provided context.
func (p *OverrideQuotaForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandOverrideQuotaForOrigin, p, nil)
}

// TrackCacheStorageForOriginParams registers origin to be notified when an
// update occurs to its cache storage list.
type TrackCacheStorageForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// TrackCacheStorageForOrigin registers origin to be notified when an update
// occurs to its cache storage list.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackCacheStorageForOrigin
//
// parameters:
//   origin - Security origin.
func TrackCacheStorageForOrigin(origin string) *TrackCacheStorageForOriginParams {
	return &TrackCacheStorageForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.trackCacheStorageForOrigin against the provided context.
func (p *TrackCacheStorageForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrackCacheStorageForOrigin, p, nil)
}

// TrackIndexedDBForOriginParams registers origin to be notified when an
// update occurs to its IndexedDB.
type TrackIndexedDBForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// TrackIndexedDBForOrigin registers origin to be notified when an update
// occurs to its IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackIndexedDBForOrigin
//
// parameters:
//   origin - Security origin.
func TrackIndexedDBForOrigin(origin string) *TrackIndexedDBForOriginParams {
	return &TrackIndexedDBForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.trackIndexedDBForOrigin against the provided context.
func (p *TrackIndexedDBForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrackIndexedDBForOrigin, p, nil)
}

// UntrackCacheStorageForOriginParams unregisters origin from receiving
// notifications for cache storage.
type UntrackCacheStorageForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// UntrackCacheStorageForOrigin unregisters origin from receiving
// notifications for cache storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackCacheStorageForOrigin
//
// parameters:
//   origin - Security origin.
func UntrackCacheStorageForOrigin(origin string) *UntrackCacheStorageForOriginParams {
	return &UntrackCacheStorageForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.untrackCacheStorageForOrigin against the provided context.
func (p *UntrackCacheStorageForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandUntrackCacheStorageForOrigin, p, nil)
}

// UntrackIndexedDBForOriginParams unregisters origin from receiving
// notifications for IndexedDB.
type UntrackIndexedDBForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// UntrackIndexedDBForOrigin unregisters origin from receiving notifications
// for IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackIndexedDBForOrigin
//
// parameters:
//   origin - Security origin.
func UntrackIndexedDBForOrigin(origin string) *UntrackIndexedDBForOriginParams {
	return &UntrackIndexedDBForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.untrackIndexedDBForOrigin against the provided context.
func (p *UntrackIndexedDBForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandUntrackIndexedDBForOrigin, p, nil)
}

// GetTrustTokensParams returns the number of stored Trust Tokens per issuer
// for the current browsing context.
type GetTrustTokensParams struct{}

// GetTrustTokens returns the number of stored Trust Tokens per issuer for
// the current browsing context.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getTrustTokens
func GetTrustTokens() *GetTrustTokensParams {
	return &GetTrustTokensParams{}
}

// GetTrustTokensReturns return values.
type GetTrustTokensReturns struct {
	Tokens []*TrustTokens `json:"tokens,omitempty"`
}

// Do executes Storage.getTrustTokens against the provided context.
//
// returns:
//   tokens
func (p *GetTrustTokensParams) Do(ctx context.Context) (tokens []*TrustTokens, err error) {
	// execute
	var res GetTrustTokensReturns
	err = cdp.Execute(ctx, CommandGetTrustTokens, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Tokens, nil
}

// ClearTrustTokensParams removes all Trust Tokens issued by the provided
// issuerOrigin. Leaves other stored data, including the issuer's Redemption
// Records, intact.
type ClearTrustTokensParams struct {
	IssuerOrigin string `json:"issuerOrigin"`
}

// ClearTrustTokens removes all Trust Tokens issued by the provided
// issuerOrigin. Leaves other stored data, including the issuer's Redemption
// Records, intact.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearTrustTokens
//
// parameters:
//   issuerOrigin
func ClearTrustTokens(issuerOrigin string) *ClearTrustTokensParams {
	return &ClearTrustTokensParams{
		IssuerOrigin: issuerOrigin,
	}
}

// ClearTrustTokensReturns return values.
type ClearTrustTokensReturns struct {
	DidDeleteTokens bool `json:"didDeleteTokens,omitempty"` // True if any tokens were deleted, false otherwise.
}

// Do executes Storage.clearTrustTokens against the provided context.
//
// returns:
//   didDeleteTokens - True if any tokens were deleted, false otherwise.
func (p *ClearTrustTokensParams) Do(ctx context.Context) (didDeleteTokens bool, err error) {
	// execute
	var res ClearTrustTokensReturns
	err = cdp.Execute(ctx, CommandClearTrustTokens, p, &res)
	if err != nil {
		return false, err
	}

	return res.DidDeleteTokens, nil
}

// Command names.
const (
	CommandClearDataForOrigin           = "Storage.clearDataForOrigin"
	CommandGetCookies                   = "Storage.getCookies"
	CommandSetCookies                   = "Storage.setCookies"
	CommandClearCookies                 = "Storage.clearCookies"
	CommandGetUsageAndQuota             = "Storage.getUsageAndQuota"
	CommandOverrideQuotaForOrigin       = "Storage.overrideQuotaForOrigin"
	CommandTrackCacheStorageForOrigin   = "Storage.trackCacheStorageForOrigin"
	CommandTrackIndexedDBForOrigin      = "Storage.trackIndexedDBForOrigin"
	CommandUntrackCacheStorageForOrigin = "Storage.untrackCacheStorageForOrigin"
	CommandUntrackIndexedDBForOrigin    = "Storage.untrackIndexedDBForOrigin"
	CommandGetTrustTokens               = "Storage.getTrustTokens"
	CommandClearTrustTokens             = "Storage.clearTrustTokens"
)
