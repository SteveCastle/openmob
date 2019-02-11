import React from 'react';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledButton = styled.button`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  outline: none;
`;

export default function Button({ onClick, label, block, dark }) {
  return (
    <StyledButton
      block={block}
      border={0}
      borderRadius={2}
      onClick={onClick}
      p={3}
      bg={dark ? 'buttons.darkBackground' : 'buttons.lightBackground'}
      fontSize={1}
      color={dark ? 'buttons.lightText' : 'buttons.darkText'}
      width={block ? 1 : null}
    >
      {label}
    </StyledButton>
  );
}
