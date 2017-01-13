package roadmap

import (
	"time"
)

// Roadmap represents a roadmap (product)
type Roadmap struct {
	ID             string       `json:"id"`
	AccountID      string       `json:"accountId,omitempty"`
	Name           string       `json:"name,omitempty"`
	Website        string       `json:"website"`
	SubDomain      string       `json:"subDomain"`
	Domain         string       `json:"domain"`
	AccentColor    string       `json:"color"`
	SecondaryColor string       `json:"altColor"`
	CSS            string       `json:"css"`
	Header         string       `json:"header"`
	Footer         string       `json:"footer"`
	Title          string       `json:"title"`
	Description    string       `json:"desc"`
	Logo           string       `json:"logo"`
	FavIcon        string       `json:"favicon"`
	MailFrom       string       `json:"from,omitempty"`
	MailReplyTo    string       `json:"replyTo,omitempty"`
	MailSubject    string       `json:"subject,omitempty"`
	MailBody       string       `json:"body,omitempty"`
	Analytics      string       `json:"analytics,omitempty"`
	NewUntilDays   int          `json:"newDays"`
	Categories     []string     `json:"categories,omitempty"`
	Teams          []string     `json:"teams,omitempty"`
	IsPrivate      bool         `json:"private"`
	VisualOptions  VisualOption `json:"visualOptions,omitempty"`
	WidgetOptions  WidgetOption `json:"widget"`
	Created        time.Time    `json:"created"`
}

// WidgetOption represents config for the widget
type WidgetOption struct {
	UIType              int    `json:"type"`
	Text                string `json:"text"`
	TitlePlaceholder    string `json:"title"`
	FeedbackPlaceholder string `json:"feedback"`
	SubmitButton        string `json:"submit"`
	ThankYou            string `json:"thanks"`
	Theme               string `json:"theme"`
}

// VisualOption represents how visual roadmap are displayed
type VisualOption struct {
	GroupByTeam bool `json:"groupByTeam"`
	Timeframe   int  `json:"timeFrame"`
}

const (
	// ItemTypeIdea represents a user feedback or team idea
	ItemTypeIdea = iota
	// ItemTypeStory represents a roadmap story
	ItemTypeStory
	// ItemTypeCompleted represents a completed item
	ItemTypeCompleted
	// ItemTypeMilestone represents a story milestone kept for reference
	ItemTypeMilestone
)

const (
	// ColumnCurrent represents current roadmap items
	ColumnCurrent = 2
	// ColumnSoon represents soon to be worked roadmap items
	ColumnSoon = 1
	// ColumnFuture represents on radar roadmap items
	ColumnFuture = 0
	// ColumnTeam represents ideas
	ColumnTeam = 2
	// ColumnUser represent user feedback
	ColumnUser = 1
	// ColumnWidget represents ideas displayed on widget
	ColumnWidget = 0
	// ColumnCompleted represents a completed item
	ColumnCompleted = 99
)

// BaseItem contains shared attributes for feedback, idea and story
type BaseItem struct {
	ID               string       `json:"id"`
	RoadmapID        string       `json:"roadmapId"`
	ItemType         int          `json:"type"`
	ColumnIndex      int          `json:"column"`
	Title            string       `json:"title"`
	Description      string       `json:"desc"`
	Hidden           bool         `json:"hidden"`
	Order            int          `json:"order"`
	Subscribers      []Subscriber `json:"subscribers"`
	Token            string       `bson:"json:"token"`
	SubscribersCount int          `json:"subscribersCount"`
	RevenueTotal     int          `json:"revenue"`
	Owner            Owner        `json:"owner"`
	Tags             []string     `json:"tags"`
	Created          time.Time    `json:"created"`
	Updated          time.Time    `json:"updated"`
	Deleted          time.Time    `json:"deleted"`
	IsDeleted        bool         `json:"isDeleted"`
}

// UserFeedback represents user feedback
type UserFeedback struct {
	BaseItem
	Comments []Comment `json:"comments"`
}

// Idea represents an idea
type Idea struct {
	UserFeedback
	Like      int      `json:"like"`
	Value     int      `json:"value"`
	Effort    int      `json:"effort"`
	Mockups   []Mockup `json:"mockups"`
	Category  string   `json:"category"`
	Team      string   `json:"team"`
	Published bool     `json:"published"`
}

// Story represents a roadmap story
type Story struct {
	Idea
	Clap        int       `json:"clap"`
	Tasks       []Task    `json:"tasks"`
	Completed   bool      `json:"completed"`
	CompletedOn time.Time `json:"completedOn"`
}

// Item to stay temporarily backward compatible
type Item struct {
	Story
}

// Owner represents the owner of an item
type Owner struct {
	FirstName       string `json:"first"`
	LastName        string `json:"last"`
	Avatar          string `json:"avatar"`
	Via             string `json:"via"`
	ViaRef          string `json:"link"`
	OriginalRequest string `json:"original"`
}

// Subscriber represents people that will be notified on completion
type Subscriber struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first"`
	LastName  string    `json:"last"`
	Email     string    `json:"email"`
	Revenue   int       `json:"revenue"`
	Created   time.Time `json:"created"`
}

const (
	// CommentTypeActivity represents a system activity comment
	CommentTypeActivity = iota
	// CommentTypeMessage represents a team member comment
	CommentTypeMessage
	// CommentTypeLike represents a team member liking an idea
	CommentTypeLike
)

// Comment represents a comment left by team member on an item
type Comment struct {
	ID       string    `json:"id"`
	Type     int       `json:"type"`
	Name     string    `json:"name"`
	Gravatar string    `json:"gravatar"`
	Body     string    `json:"body"`
	Created  time.Time `json:"created"`
}

// Task represents one action to complete for a roadmap item
type Task struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Done            bool      `json:"done"`
	AssignTo        string    `json:"assign"`
	Gravatar        string    `json:"g"`
	Created         time.Time `json:"created"`
	Completed       time.Time `json:"completed"`
	TrelloID        string    `json:"trelloId"`
	GitHubID        int       `json:"ghId"`
	OriginalIdeaID  *string   `json:"ideaId"`
	SubscriberCount int       `json:"subCount"`
	RevenueTotal    int       `json:"revenue"`
}

// Mockup represents a mockup for a roadmap item
type Mockup struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	IsFeatured bool   `json:"featured"`
}
