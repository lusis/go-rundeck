package responses

// UserProfileResponse represents a user info response
// http://rundeck.org/docs/api/index.html#get-user-profile
type UserProfileResponse struct {
	Login     string `json:"login"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

func (u UserProfileResponse) minVersion() int  { return 21 }
func (u UserProfileResponse) maxVersion() int  { return CurrentVersion }
func (u UserProfileResponse) deprecated() bool { return false }

// UserProfileResponseTestFile is test data for a UserInfoResponse
const UserProfileResponseTestFile = "user.json"

// ListUsersResponse is a collection of `UserInfo`
// http://rundeck.org/docs/api/index.html#list-users
type ListUsersResponse []UserProfileResponse

// ListUsersResponseTestFile is test data for a UsersInfoResponse
const ListUsersResponseTestFile = "users.json"

func (u ListUsersResponse) minVersion() int  { return 21 }
func (u ListUsersResponse) maxVersion() int  { return CurrentVersion }
func (u ListUsersResponse) deprecated() bool { return false }
