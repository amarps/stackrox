package userpki

import (
	"time"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/authproviders"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/auth/permissions/utils"
	"github.com/stackrox/rox/pkg/grpc/authn"
	"github.com/stackrox/rox/pkg/grpc/requestinfo"
)

var _ authn.Identity = (*identity)(nil)

type identity struct {
	info          requestinfo.CertInfo
	provider      authproviders.Provider
	resolvedRoles []permissions.ResolvedRole
	attributes    map[string][]string
}

func (i *identity) Attributes() map[string][]string {
	return i.attributes
}

func (i *identity) FriendlyName() string {
	return i.info.Subject.CommonName
}

func (i *identity) FullName() string {
	return i.info.Subject.CommonName
}

func (i *identity) User() *storage.UserInfo {
	return &storage.UserInfo{
		FriendlyName: i.info.Subject.CommonName,
		Permissions:  i.Permissions(),
		Roles:        utils.ExtractRolesForUserInfo(i.resolvedRoles),
	}
}

func (i *identity) Permissions() *storage.ResourceToAccess {
	return utils.NewUnionPermissions(i.resolvedRoles)
}

func (i *identity) Roles() []permissions.ResolvedRole {
	return i.resolvedRoles
}

func (i *identity) Service() *storage.ServiceIdentity {
	return nil
}

func (i *identity) Expiry() time.Time {
	return i.info.NotAfter
}

func (i *identity) ExternalAuthProvider() authproviders.Provider {
	return i.provider
}

func (i *identity) UID() string {
	return userID(i.info)
}

func userID(info requestinfo.CertInfo) string {
	return "userpki:" + info.CertFingerprint
}