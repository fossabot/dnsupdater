// Package constants contains all constants needed in this programme
package constants

const MsgHeaderDNSRecordUpdateSuccessful = "Successfully updated the DNS record"
const MsgHeaderCurrentIPAddr = "Current IP address is:"
const MsgHeaderFetchingIPOfDomain = "Fetching the IP address of domain"
const MsgHeaderLoadingConfig = "Loading configuration from"

const MsgFormatDNSFetchResult = "The IP address of %s is %s."
const MsgFormatUpdatingDNS = "Updating the IP address of domain %s to %s."

const ErrMsgHeaderFetchDomainInfoFailed = "Failed to get the information for domain"
const ErrMsgHeaderUpdateDNSRecordFailed = "Failed to update the DNS record"

const MsgIPAddrNotChanged = "IP address not changed. Will not update the DNS record. "
const MsgCheckingCurrentIPAddr = "Checking current IP address..."

const ErrDomainNameNotExist = "Domain name does not exist. "
const ErrCloseHTTPConnectionFail = "Error closing the HTTP connection. "
