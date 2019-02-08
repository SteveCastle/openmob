import React from 'react';
import styled from 'styled-components';

const StyledButton = styled.button`
  background-color: #000000;
  border: none;
  border-radius: 4px;
  color: #a5aed5;
  font-size: 1rem;
  padding: 16px;
`;

export default function Button({ label }) {
  return <StyledButton>{label}</StyledButton>;
}
