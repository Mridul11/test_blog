package actions

import (
	"fmt"
	"net/http"
	"test_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Widget)
// DB Table: Plural (widgets)
// Resource: Plural (Widgets)
// Path: Plural (/widgets)
// View Template Folder: Plural (/templates/widgets/)

// WidgetsResource is the resource for the Widget model
type WidgetsResource struct {
	buffalo.Resource
}

// List gets all Widgets. This function is mapped to the path
// GET /widgets
func (v WidgetsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	widgets := &models.Widgets{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Widgets from the DB
	if err := q.All(widgets); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("widgets", widgets)
		return c.Render(http.StatusOK, r.HTML("/widgets/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(widgets))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(widgets))
	}).Respond(c)
}

// Show gets the data for one Widget. This function is mapped to
// the path GET /widgets/{widget_id}
func (v WidgetsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("widget", widget)

		return c.Render(http.StatusOK, r.HTML("/widgets/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(widget))
	}).Respond(c)
}

// New renders the form for creating a new Widget.
// This function is mapped to the path GET /widgets/new
func (v WidgetsResource) New(c buffalo.Context) error {
	c.Set("widget", &models.Widget{})

	return c.Render(http.StatusOK, r.HTML("/widgets/new.plush.html"))
}

// Create adds a Widget to the DB. This function is mapped to the
// path POST /widgets
func (v WidgetsResource) Create(c buffalo.Context) error {
	// Allocate an empty Widget
	widget := &models.Widget{}

	// Bind widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(widget)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("widget", widget)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/widgets/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "widget.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/widgets/%v", widget.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(widget))
	}).Respond(c)
}

// Edit renders a edit form for a Widget. This function is
// mapped to the path GET /widgets/{widget_id}/edit
func (v WidgetsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("widget", widget)
	return c.Render(http.StatusOK, r.HTML("/widgets/edit.plush.html"))
}

// Update changes a Widget in the DB. This function is mapped to
// the path PUT /widgets/{widget_id}
func (v WidgetsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Widget to the html form elements
	if err := c.Bind(widget); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(widget)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("widget", widget)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/widgets/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "widget.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/widgets/%v", widget.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(widget))
	}).Respond(c)
}

// Destroy deletes a Widget from the DB. This function is mapped
// to the path DELETE /widgets/{widget_id}
func (v WidgetsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Widget
	widget := &models.Widget{}

	// To find the Widget the parameter widget_id is used.
	if err := tx.Find(widget, c.Param("widget_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(widget); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "widget.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/widgets")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(widget))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(widget))
	}).Respond(c)
}
