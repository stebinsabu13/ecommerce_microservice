package repository

import (
	"context"
	"errors"

	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/domain"
	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/pb"
	"github.com/stebinsabu13/ecommerce_microservice/product_service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type ProductDatabase struct {
	DB *gorm.DB
}

func NewProductRepo(db *gorm.DB) interfaces.ProductRepo {
	return &ProductDatabase{
		DB: db,
	}
}

func (c *ProductDatabase) FindAllProducts(ctx context.Context, req *pb.ListAllProductRequest) ([]*pb.Product, error) {
	var products []*pb.Product
	query := `SELECT p.id,p.model_name,p.image,b.brand_name,c.category_name FROM products p
	LEFT JOIN brands b on p.brand_id=b.id
	INNER JOIN categories c on p.category_id=c.id where p.deleted_at is null LIMIT $1 OFFSET $2`
	result := c.DB.Raw(query, req.Limit, req.Page).Scan(&products)
	if result.Error != nil {
		return products, errors.New("failed to load products")
	}
	return products, nil
}

func (c *ProductDatabase) AddProduct(ctx context.Context, product domain.Product) error {
	result := c.DB.Create(&product).Error
	if result != nil {
		return errors.New("failed to add product")
	}
	return nil
}

func (c *ProductDatabase) FindProductById(ctx context.Context, req *pb.ListProductDetailsRequest) ([]*pb.ProdDetail, error) {
	var Product []*pb.ProdDetail
	//Finding the Product
	query := `select p.model_name,p.image,b.brand_name,pd.stock,pd.price,c.colour,s.size,d.percentage from products p
	left join brands b on b.id=p.brand_id
	inner join product_details pd on pd.product_id=p.id
	inner join available_colours c on c.id=pd.available_colour_id
	inner join available_sizes s on s.id=pd.available_size_id
	left join discounts d on d.id=pd.discount_id where p.id=$1 and pd.deleted_at is null LIMIT $2 OFFSET $3`
	result := c.DB.Raw(query, req.Productid, req.Limit, req.Page).Scan(&Product)
	if result.Error != nil {
		return Product, errors.New("failed to get product")
	}
	return Product, nil
}

func (c *ProductDatabase) AddProductDetail(ctx context.Context, productdetail domain.ProductDetails) error {
	result := c.DB.Create(&productdetail).Error
	if result != nil {
		return result
	}
	return nil
}
