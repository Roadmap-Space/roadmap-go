package roadmap_test

func deleteItem(id, token string) error {
  return c.Items.Delete(id, token)
}