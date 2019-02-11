import React from 'react';
import styled from 'styled-components';
import { space, width, fontSize, color } from 'styled-system';

const StyledButton = styled.button`
  border: none;
  border-radius: 4px;
  ${space}
  ${width}
  ${fontSize}
  ${color}
  ${props => (props.block ? 'width: 100%;' : null)}
`;

export default function Button({ onClick, label, block }) {
  return (
    <StyledButton
      block={block}
      onClick={onClick}
      p={3}
      bg="buttons.darkBackground"
      color="buttons.lightText"
    >
      {label}
    </StyledButton>
  );
}
