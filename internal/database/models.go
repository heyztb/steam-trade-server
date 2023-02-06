package database

type Bot struct {
	ID             int64
	Username       string
	Passwd         string
	SharedSecret   string
	IdentitySecret string
}

type Item struct {
	ID         int64
	BotID      int64
	AppID      int64
	AssetID    int64
	ClassID    int64
	InstanceID int64
}
