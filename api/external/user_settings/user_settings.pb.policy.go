// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/user_settings/user_settings.proto

package user_settings

import policy "github.com/chef/automate/api/external/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.user_settings.UserSettingsService/GetUserSettings", "global:userSettings", "global:userSettings:get", "GET", "/api/v0/user/{id}/settings", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*GetUserSettingsRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.user_settings.UserSettingsService/UpdateUserSettings", "global:userSettings", "global:userSettings:update", "PUT", "/api/v0/user/{id}/settings", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*UpdateUserSettingsRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.user_settings.UserSettingsService/DeleteUserSettings", "global:userSettings", "global:userSettings:delete", "DELETE", "/api/v0/user/{id}/settings", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*DeleteUserSettingsRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
}
