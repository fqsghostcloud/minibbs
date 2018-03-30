package models

type RolePermissions struct {
	id            int
	role_id       int
	permission_id int
}

type TopicTags struct {
	id       int
	topic_id int
	tag_id   int
}
