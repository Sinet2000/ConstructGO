package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	ChildCategories []Category         `json:"childCategories,omitempty"`
	Items           []Item             `json:"items,omitempty"`
}

func (c *Category) AddChildCategory(child Category) {
	c.ChildCategories = append(c.ChildCategories, child)
}

func (c *Category) HasChildCategory(child Category) bool {
	for _, ch := range c.ChildCategories {
		if ch.ID == child.ID {
			return true
		}
	}
	return false
}

func (c *Category) AddChildCategories(children []Category) {
	c.ChildCategories = append(c.ChildCategories, children...)
}

func (c *Category) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c *Category) HasItem(item Item) bool {
	for _, i := range c.Items {
		if i.ID == item.ID {
			return true
		}
	}
	return false
}

func (c *Category) AddItems(items []Item) {
	c.Items = append(c.Items, items...)
}

func (c *Category) RemoveChildCategory(childID primitive.ObjectID) error {
	for i, ch := range c.ChildCategories {
		if ch.ID == childID {
			c.ChildCategories = append(c.ChildCategories[:i], c.ChildCategories[i+1:]...)
			return nil
		}
	}
	return errors.New("child category not found")
}

func (c *Category) RemoveItem(itemID primitive.ObjectID) error {
	for i, item := range c.Items {
		if item.ID == itemID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func (c *Category) UpdateItem(updatedItem Item) error {
	for i, item := range c.Items {
		if item.ID == updatedItem.ID {
			c.Items[i] = updatedItem
			return nil
		}
	}
	return errors.New("item not found")
}

func (c *Category) GetItemByID(itemID primitive.ObjectID) (*Item, error) {
	for _, item := range c.Items {
		if item.ID == itemID {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (c *Category) GetChildCategoryByID(childID primitive.ObjectID) (*Category, error) {
	for _, ch := range c.ChildCategories {
		if ch.ID == childID {
			return &ch, nil
		}
	}
	return nil, errors.New("child category not found")
}
