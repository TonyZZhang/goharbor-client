// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Configurations configurations
//
// swagger:model Configurations
type Configurations struct {

	// The auth mode of current system, such as "db_auth", "ldap_auth"
	AuthMode string `json:"auth_mode,omitempty"`

	// The default count quota for the new created projects.
	CountPerProject string `json:"count_per_project,omitempty"`

	// The sender name for Email notification.
	EmailFrom string `json:"email_from,omitempty"`

	// The hostname of SMTP server that sends Email notification.
	EmailHost string `json:"email_host,omitempty"`

	// By default it's empty so the email_username is picked.
	EmailIdentity string `json:"email_identity,omitempty"`

	// Whether or not the certificate will be verified when Harbor tries to access the email server.
	EmailInsecure bool `json:"email_insecure,omitempty"`

	// The port of SMTP server.
	EmailPort int64 `json:"email_port,omitempty"`

	// When it's set to true the system will access Email server via TLS by default.  If it's set to false, it still will handle "STARTTLS" from server side.
	EmailSsl bool `json:"email_ssl,omitempty"`

	// The username for authenticate against SMTP server.
	EmailUsername string `json:"email_username,omitempty"`

	// The Base DN for LDAP binding.
	LdapBaseDn string `json:"ldap_base_dn,omitempty"`

	// The filter for LDAP binding.
	LdapFilter string `json:"ldap_filter,omitempty"`

	// Specify the ldap group which have the same privilege with Harbor admin.
	LdapGroupAdminDn string `json:"ldap_group_admin_dn,omitempty"`

	// The attribute which is used as identity of the LDAP group, default is cn.
	LdapGroupAttributeName string `json:"ldap_group_attribute_name,omitempty"`

	// The base DN to search LDAP group.
	LdapGroupBaseDn string `json:"ldap_group_base_dn,omitempty"`

	// The filter to search the ldap group.
	LdapGroupSearchFilter string `json:"ldap_group_search_filter,omitempty"`

	// The scope to search ldap. '0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE'
	LdapGroupSearchScope int64 `json:"ldap_group_search_scope,omitempty"`

	// 0-LDAP_SCOPE_BASE, 1-LDAP_SCOPE_ONELEVEL, 2-LDAP_SCOPE_SUBTREE
	LdapScope int64 `json:"ldap_scope,omitempty"`

	// The DN of the user to do the search.
	LdapSearchDn string `json:"ldap_search_dn,omitempty"`

	// timeout in seconds for connection to LDAP server.
	LdapTimeout int64 `json:"ldap_timeout,omitempty"`

	// The attribute which is used as identity for the LDAP binding, such as "CN" or "SAMAccountname"
	LdapUID string `json:"ldap_uid,omitempty"`

	// The URL of LDAP server.
	LdapURL string `json:"ldap_url,omitempty"`

	// The client id of the OIDC.
	OidcClientID string `json:"oidc_client_id,omitempty"`

	// The client secret of the OIDC.
	OidcClientSecret string `json:"oidc_client_secret,omitempty"`

	// The URL of an OIDC-complaint server, must start with 'https://'.
	OidcEndpoint string `json:"oidc_endpoint,omitempty"`

	// The name of the OIDC provider.
	OidcName string `json:"oidc_name,omitempty"`

	// The scope sent to OIDC server during authentication, should be separated by comma. It has to contain “openid”, and “offline_access”. If you are using google, please remove “offline_access” from this field.
	OidcScope string `json:"oidc_scope,omitempty"`

	// Whether verify your OIDC server certificate, disable it if your OIDC server is hosted via self-hosted certificate.
	OidcVerifyCert bool `json:"oidc_verify_cert,omitempty"`

	// This attribute restricts what users have the permission to create project.  It can be "everyone" or "adminonly".
	ProjectCreationRestriction string `json:"project_creation_restriction,omitempty"`

	// This attribute indicates whether quota per project enabled in harbor
	QuotaPerProjectEnable bool `json:"quota_per_project_enable,omitempty"`

	// 'docker push' is prohibited by Harbor if you set it to true.
	ReadOnly bool `json:"read_only,omitempty"`

	// scan all policy
	ScanAllPolicy *ConfigurationsScanAllPolicy `json:"scan_all_policy,omitempty"`

	// Whether the Harbor instance supports self-registration.  If it's set to false, admin need to add user to the instance.
	SelfRegistration bool `json:"self_registration,omitempty"`

	// The default storage quota for the new created projects.
	StoragePerProject string `json:"storage_per_project,omitempty"`

	// The expiration time of the token for internal Registry, in minutes.
	TokenExpiration int64 `json:"token_expiration,omitempty"`

	// Whether or not the certificate will be verified when Harbor tries to access a remote Harbor instance for replication.
	VerifyRemoteCert bool `json:"verify_remote_cert,omitempty"`
}

// Validate validates this configurations
func (m *Configurations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScanAllPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configurations) validateScanAllPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ScanAllPolicy) { // not required
		return nil
	}

	if m.ScanAllPolicy != nil {
		if err := m.ScanAllPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scan_all_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configurations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configurations) UnmarshalBinary(b []byte) error {
	var res Configurations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConfigurationsScanAllPolicy configurations scan all policy
//
// swagger:model ConfigurationsScanAllPolicy
type ConfigurationsScanAllPolicy struct {

	// parameter
	Parameter *ConfigurationsScanAllPolicyParameter `json:"parameter,omitempty"`

	// The type of scan all policy, currently the valid values are "none" and "daily"
	Type string `json:"type,omitempty"`
}

// Validate validates this configurations scan all policy
func (m *ConfigurationsScanAllPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigurationsScanAllPolicy) validateParameter(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameter) { // not required
		return nil
	}

	if m.Parameter != nil {
		if err := m.Parameter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scan_all_policy" + "." + "parameter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationsScanAllPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationsScanAllPolicy) UnmarshalBinary(b []byte) error {
	var res ConfigurationsScanAllPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConfigurationsScanAllPolicyParameter The parameters of the policy, the values are dependant on the type of the policy.
//
// swagger:model ConfigurationsScanAllPolicyParameter
type ConfigurationsScanAllPolicyParameter struct {

	// The offset in seconds of UTC 0 o'clock, only valid when the policy type is "daily"
	DailyTime int64 `json:"daily_time,omitempty"`
}

// Validate validates this configurations scan all policy parameter
func (m *ConfigurationsScanAllPolicyParameter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigurationsScanAllPolicyParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigurationsScanAllPolicyParameter) UnmarshalBinary(b []byte) error {
	var res ConfigurationsScanAllPolicyParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
