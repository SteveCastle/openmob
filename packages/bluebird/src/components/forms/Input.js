import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledInput = styled.input`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  outline: none;
`;

function Input({ onClick, label, block, dark }) {
  return (
    <StyledInput
      block={block}
      border={0}
      onClick={onClick}
      p={3}
      bg={dark ? 'forms.darkBackground' : 'forms.lightBackground'}
      fontSize={1}
      color={dark ? 'forms.lightText' : 'forms.darkText'}
      width={block ? 1 : null}
    >
      {label}
    </StyledInput>
  );
}

Input.propTypes = {
  onClick: PropTypes.func,
  label: PropTypes.string,
  icon: PropTypes.element,
  loader: PropTypes.element,
  state: PropTypes.oneOf(['ready', 'valid', 'error']),
  block: PropTypes.bool,
  dark: PropTypes.bool
};
export default Input;