package integration_tests

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	. "gorm.io/gorm"
	"split-the-bill-server/domain/model"
	. "split-the-bill-server/storage/database/entity"
	"time"
)

type Seed struct {
	Name string
	Run  func(*DB) error
}

var (
	// USERS
	User1 = User{
		Base:  Base{ID: uuid.New()},
		Email: "felix@gmail.com",
	}

	User2 = User{
		Base:  Base{ID: uuid.New()},
		Email: "marivn@gmail.com",
	}

	User3 = User{
		Base:  Base{ID: uuid.New()},
		Email: "jan@gmail.com",
	}

	// COOKIES
	Cookie1 = AuthCookie{
		Base:        Base{ID: uuid.New()},
		UserID:      User1.ID,
		ValidBefore: time.Now().Add(model.SessionCookieValidityPeriod),
	}

	// PASSWORD
	Password = "test"

	// CREDENTIALS
	Pw, _ = bcrypt.GenerateFromPassword([]byte("test"), 10)

	Credentials1 = Credentials{
		UserID: User1.ID,
		Hash:   Pw,
	}

	Credentials2 = Credentials{
		UserID: User2.ID,
		Hash:   Pw,
	}

	Credentials3 = Credentials{
		UserID: User3.ID,
		Hash:   Pw,
	}

	// GROUPS
	Group1 = Group{
		Base:     Base{ID: uuid.New()},
		Name:     "Wohnung",
		OwnerUID: User1.ID,
		Members:  []*User{&User1, &User2, &User3},
	}

	Group2 = Group{
		Base:     Base{ID: uuid.New()},
		Name:     "Urlaub",
		OwnerUID: User1.ID,
		Members:  []*User{&User1, &User2},
	}

	Group3 = Group{
		Base:     Base{ID: uuid.New()},
		Name:     "Partyyy",
		OwnerUID: User2.ID,
		Members:  []*User{&User2, &User3},
	}

	// ITEMS
	Item1 = Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Bread",
		Price:        2.5,
		Contributors: []*User{&User1, &User2},
	}

	Item2 = Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Milk",
		Price:        1.5,
		Contributors: []*User{&User1},
	}

	Item3 = Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Miete",
		Price:        1050,
		Contributors: []*User{&User1, &User2, &User3},
	}

	Item4 = Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Sunscreen",
		Price:        10.0,
		Contributors: []*User{&User1, &User2},
	}

	Item5 = Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Beach Towels",
		Price:        15.0,
		Contributors: []*User{&User1, &User2},
	}

	Item6 = Item{
		Base:         Base{ID: uuid.New()},
		Name:         "Snorkel Gear",
		Price:        25.0,
		Contributors: []*User{&User2},
	}

	// BILLS
	Bill1 = Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: User1.ID,
		Name:    "Groceries",
		Date:    time.Now(),
		Items:   []Item{Item1, Item2},
		GroupID: Group1.ID,
	}

	Bill2 = Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: User1.ID,
		Name:    "Miete",
		Date:    time.Now(),
		Items:   []Item{Item3},
		GroupID: Group1.ID,
	}

	Bill3 = Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: User1.ID,
		Name:    "Beach Trip Expenses",
		Date:    time.Now().AddDate(0, 0, 10),
		Items:   []Item{Item4, Item5},
		GroupID: Group2.ID,
	}

	Bill4 = Bill{
		Base:    Base{ID: uuid.New()},
		OwnerID: User2.ID,
		Name:    "Water Sports",
		Date:    time.Now().AddDate(0, 0, 15),
		Items:   []Item{Item6},
		GroupID: Group2.ID,
	}

	// INVITATIONS
	Invitation1 = GroupInvitation{
		Base:    Base{ID: uuid.New()},
		GroupID: Group1.ID,
	}

	Invitation2 = GroupInvitation{
		Base:    Base{ID: uuid.New()},
		GroupID: Group2.ID,
	}

	Invitation3 = GroupInvitation{
		Base:    Base{ID: uuid.New()},
		GroupID: Group3.ID,
	}
)

func All() []Seed {

	return []Seed{
		{
			Name: "CreateUsers",
			Run: func(db *DB) error {
				if err := db.Create(&User1).Error; err != nil {
					return err
				}
				if err := db.Create(&User2).Error; err != nil {
					return err
				}
				if err := db.Create(&User3).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateCredentials",
			Run: func(db *DB) error {
				if err := db.Create(&Credentials1).Error; err != nil {
					return err
				}
				if err := db.Create(&Credentials2).Error; err != nil {
					return err
				}
				if err := db.Create(&Credentials3).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateAuthCookies",
			Run: func(db *DB) error {
				if err := db.Create(&Cookie1).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateGroups",
			Run: func(db *DB) error {
				if err := db.Create(&Group1).Error; err != nil {
					return err
				}
				if err := db.Create(&Group2).Error; err != nil {
					return err
				}
				if err := db.Create(&Group3).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateBills",
			Run: func(db *DB) error {
				if err := db.Create(&Bill1).Error; err != nil {
					return err
				}
				if err := db.Create(&Bill2).Error; err != nil {
					return err
				}
				if err := db.Create(&Bill3).Error; err != nil {
					return err
				}
				if err := db.Create(&Bill4).Error; err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "CreateGroupInvitations",
			Run: func(db *DB) error {
				if err := db.Create(&Invitation1).Error; err != nil {
					return err
				}
				if err := db.Create(&Invitation2).Error; err != nil {
					return err
				}
				if err := db.Create(&Invitation3).Error; err != nil {
					return err
				}
				return nil
			},
		},
	}
}
