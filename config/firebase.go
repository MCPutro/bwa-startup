package config

type FirebaseConfig interface {
	ProjectId() string
	ServiceAccount() string
	AdminKey() string
	BucketName() string
	BucketPath() string
}

type Firebase struct {
	FirebaseProjectId        string `mapstructure:"project-id"`
	FirebaseServiceAccountId string `mapstructure:"service-account-id"`
	FirebaseAdminKey         string `mapstructure:"admin-key"`
	FirebaseBucket           Bucket `mapstructure:"bucket"`
}

type Bucket struct {
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}

func (f *Firebase) ProjectId() string {
	return f.FirebaseProjectId
}

func (f *Firebase) ServiceAccount() string {
	return f.FirebaseServiceAccountId
}

func (f *Firebase) AdminKey() string {
	return f.FirebaseAdminKey
}

func (f *Firebase) BucketName() string {
	return f.FirebaseBucket.Name
}

func (f *Firebase) BucketPath() string {
	return f.FirebaseBucket.Path
}
