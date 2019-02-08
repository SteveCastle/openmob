import React from 'react';
import styled from 'styled-components';

const StyledButton = styled.button`
  background-color: ${props => props.theme.main};
  border: none;
  border-radius: 4px;
  color: ${props => props.theme.text};
  display: block;
  font-size: 1rem;
  padding: 16px;
  ${props => (props.block ? 'width: 100%;' : null)}
`;

export default function Button({ onClick, label, block }) {
  return (
    <StyledButton block={block} onClick={onClick}>
      {label}
    </StyledButton>
  );
}
