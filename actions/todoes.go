package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
	"net/http"
	"test_app/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Todo)
// DB Table: Plural (todoes)
// Resource: Plural (Todoes)
// Path: Plural (/todoes)
// View Template Folder: Plural (/templates/todoes/)

// TodoesResource is the resource for the Todo model
type TodoesResource struct {
	buffalo.Resource
}

// List gets all Todoes. This function is mapped to the path
// GET /todoes
func (v TodoesResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	todoes := &models.Todoes{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Todoes from the DB
	if err := q.All(todoes); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("todoes", todoes)
		return c.Render(http.StatusOK, r.HTML("/todoes/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(todoes))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(todoes))
	}).Respond(c)
}

// Show gets the data for one Todo. This function is mapped to
// the path GET /todoes/{todo_id}
func (v TodoesResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	// To find the Todo the parameter todo_id is used.
	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("todo", todo)

		return c.Render(http.StatusOK, r.HTML("/todoes/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(todo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(todo))
	}).Respond(c)
}

// New renders the form for creating a new Todo.
// This function is mapped to the path GET /todoes/new
func (v TodoesResource) New(c buffalo.Context) error {
	c.Set("todo", &models.Todo{})

	return c.Render(http.StatusOK, r.HTML("/todoes/new.plush.html"))
}

// Create adds a Todo to the DB. This function is mapped to the
// path POST /todoes
func (v TodoesResource) Create(c buffalo.Context) error {
	// Allocate an empty Todo
	todo := &models.Todo{}

	// Bind todo to the html form elements
	if err := c.Bind(todo); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(todo)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("todo", todo)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/todoes/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "todo.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/todoes/%v", todo.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(todo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(todo))
	}).Respond(c)
}

// Edit renders a edit form for a Todo. This function is
// mapped to the path GET /todoes/{todo_id}/edit
func (v TodoesResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("todo", todo)
	return c.Render(http.StatusOK, r.HTML("/todoes/edit.plush.html"))
}

// Update changes a Todo in the DB. This function is mapped to
// the path PUT /todoes/{todo_id}
func (v TodoesResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Todo to the html form elements
	if err := c.Bind(todo); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(todo)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("todo", todo)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("/todoes/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "todo.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/todoes/%v", todo.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(todo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(todo))
	}).Respond(c)
}

// Destroy deletes a Todo from the DB. This function is mapped
// to the path DELETE /todoes/{todo_id}
func (v TodoesResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Todo
	todo := &models.Todo{}

	// To find the Todo the parameter todo_id is used.
	if err := tx.Find(todo, c.Param("todo_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(todo); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "todo.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/todoes")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(todo))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(todo))
	}).Respond(c)
}