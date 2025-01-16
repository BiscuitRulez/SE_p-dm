import React from 'react';
import styled from 'styled-components';

interface ProductProps {
  title: string;
  price: number;
  image: string;
}

const Card = styled.div`
  border: 1px solid #ddd;
  padding: 15px;
  border-radius: 5px;
  width: 200px;
  text-align: center;
`;

const ProductCard: React.FC<ProductProps> = ({ title, price, image }) => {
  return (
    <Card>
      <img src={image} alt={title} style={{ width: '100%', height: '150px' }} />
      <h3>{title}</h3>
      <p>฿{price}</p>
      <button>เพิ่มในตะกร้า</button>
    </Card>
  );
};

export default ProductCard;
