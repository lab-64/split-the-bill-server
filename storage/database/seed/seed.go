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
		Base:  Base{ID: uuid.New()},
		Email: "felix@gmail.com",
	}

	user2 := User{
		Base:  Base{ID: uuid.New()},
		Email: "marivn@gmail.com",
	}

	user3 := User{
		Base:  Base{ID: uuid.New()},
		Email: "jan@gmail.com",
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

	// ITEMS
	item1 := Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Bread",
		Price:        2.5,
		Contributors: []*User{&user1, &user2},
	}

	item2 := Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Milk",
		Price:        1.5,
		Contributors: []*User{&user1},
	}

	item3 := Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Miete",
		Price:        1050,
		Contributors: []*User{&user1, &user2, &user3},
	}

	item4 := Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Sunscreen",
		Price:        10.0,
		Contributors: []*User{&user1, &user2},
	}

	item5 := Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Beach Towels",
		Price:        15.0,
		Contributors: []*User{&user1, &user2},
	}

	item6 := Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Snorkel Gear",
		Price:        25.0,
		Contributors: []*User{&user2},
	}

	// BILLS
	bill1 := Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: user1.ID,
		Name:    "Groceries",
		Date:    time.Now(),
		Items:   []Item{item1, item2},
		GroupID: group1.ID,
	}

	bill2 := Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: user1.ID,
		Name:    "Miete",
		Date:    time.Now(),
		Items:   []Item{item3},
		GroupID: group1.ID,
	}

	bill3 := Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: user1.ID,
		Name:    "Beach Trip Expenses",
		Date:    time.Now().AddDate(0, 0, 10),
		Items:   []Item{item4, item5},
		GroupID: group2.ID,
	}

	bill4 := Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: user2.ID,
		Name:    "Water Sports",
		Date:    time.Now().AddDate(0, 0, 15),
		Items:   []Item{item6},
		GroupID: group2.ID,
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
			Name: "CreateBills",
			Run: func(db *DB) error {
				err := db.Create(&bill1).Error
				err = db.Create(&bill2).Error
				err = db.Create(&bill3).Error
				err = db.Create(&bill4).Error
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
