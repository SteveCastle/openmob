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
  box-sizing: border-box;
`;

function Input({ onChange, label, block, dark, value, placeholder, disabled }) {
  return (
    <StyledInput
      border={0}
      onChange={onChange}
      p={3}
      m={2}
      borderRadius={1}
      bg={dark ? 'forms.darkBackground' : 'forms.lightBackground'}
      fontSize={1}
      color={dark ? 'forms.lightText' : 'forms.darkText'}
      width={block ? 1 : 2 / 3}
      value={value}
      placeholder={placeholder}
      disabled={disabled}
    >
      {label}
    </StyledInput>
  );
}

Input.propTypes = {
  onChange: PropTypes.func,
  label: PropTypes.string,
  state: PropTypes.oneOf(['ready', 'valid', 'error']),
  block: PropTypes.bool,
  dark: PropTypes.bool,
  value: PropTypes.string,
  placeholder: PropTypes.string,
  disabled: PropTypes.bool
};
export default Input;
