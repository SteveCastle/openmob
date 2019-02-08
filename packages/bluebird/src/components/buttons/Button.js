import React from 'react';
import styled from 'styled-components';

const StyledButton = styled.button`
  background-color: #000000;
  border: none;
  border-radius: 4px;
  color: #a5aed5;
  display: block;
  font-size: 1rem;
  padding: 16px;
  ${props => (props.block ? 'width: 100%;' : null)}
`;

export default function Button({ label, block = true }) {
  return <StyledButton block={block}>{label}</StyledButton>;
}
