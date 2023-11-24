package seed

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	. "gorm.io/gorm"
	. "split-the-bill-server/storage/database/entity"
	"time"
)

type Seed struct {
	Name string
	Run  func(*DB) error
}

func All() []Seed {

	// USERS
	user1 := User{
		Base:     Base{ID: uuid.New()},
		Username: "Felix",
	}

	user2 := User{
		Base:     Base{ID: uuid.New()},
		Username: "User2",
	}

	user3 := User{
		Base:     Base{ID: uuid.New()},
		Username: "User3",
	}

	// CREDENTIALS
	pw, _ := bcrypt.GenerateFromPassword([]byte("test"), 10)

	credentials1 := Credentials{
		UserID: user1.ID,
		Hash:   pw,
	}

	credentials2 := Credentials{
		UserID: user2.ID,
		Hash:   pw,
	}

	credentials3 := Credentials{
		UserID: user3.ID,
		Hash:   pw,
	}

	// GROUPS
	group1 := Group{
		Base:     Base{ID: uuid.New()},
		Name:     "Wohnung",
		OwnerUID: user1.ID,
		Members:  []*User{&user1, &user2, &user3},
	}

	group2 := Group{
		Base:     Base{ID: uuid.New()},
		Name:     "Urlaub",
		OwnerUID: user1.ID,
		Members:  []*User{&user1, &user2},
	}

	group3 := Group{
		Base:     Base{ID: uuid.New()},
		Name:     "Partyyy",
		OwnerUID: user2.ID,
		Members:  []*User{&user2, &user3},
	}

	// INVITATIONS
	invitation1 := GroupInvitation{
		Base:      Base{ID: uuid.New()},
		Date:      time.Now(),
		GroupID:   group1.ID,
		InviteeID: user2.ID,
	}

	invitation2 := GroupInvitation{
		Base:      Base{ID: uuid.New()},
		Date:      time.Now(),
		GroupID:   group2.ID,
		InviteeID: user3.ID,
	}

	invitation3 := GroupInvitation{
		Base:      Base{ID: uuid.New()},
		Date:      time.Now(),
		GroupID:   group3.ID,
		InviteeID: user1.ID,
	}

	return []Seed{
		{
			Name: "CreateUsers",
			Run: func(db *DB) error {
				err := db.Create(&user1).Error
				err = db.Create(&user2).Error
				err = db.Create(&user3).Error
				return err
			},
		},
		{
			Name: "CreateCredentials",
			Run: func(db *DB) error {
				err := db.Create(&credentials1).Error
				err = db.Create(&credentials2).Error
				err = db.Create(&credentials3).Error
				return err
			},
		},
		{
			Name: "CreateGroups",
			Run: func(db *DB) error {
				err := db.Create(&group1).Error
				err = db.Create(&group2).Error
				err = db.Create(&group3).Error
				return err
			},
		},
		{
			Name: "CreateGroupInvitations",
			Run: func(db *DB) error {
				err := db.Create(&invitation1).Error
				err = db.Create(&invitation2).Error
				err = db.Create(&invitation3).Error
				return err
			},
		},
	}
}
