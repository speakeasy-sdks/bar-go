// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Ingredient struct {
	// The name of the ingredient.
	Name string `json:"name"`
	// The type of ingredient.
	Type IngredientType `json:"type"`
	// The number of units of the ingredient in stock, only available when authenticated.
	Stock *int64 `json:"stock,omitempty"`
	// The product code of the ingredient, only available when authenticated.
	ProductCode *string `json:"productCode,omitempty"`
}

func (o *Ingredient) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Ingredient) GetType() IngredientType {
	if o == nil {
		return IngredientType("")
	}
	return o.Type
}

func (o *Ingredient) GetStock() *int64 {
	if o == nil {
		return nil
	}
	return o.Stock
}

func (o *Ingredient) GetProductCode() *string {
	if o == nil {
		return nil
	}
	return o.ProductCode
}
