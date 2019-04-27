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

const StyledTextArea = styled.textarea`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  outline: none;
  box-sizing: border-box;
`;

function TextArea({
  name,
  onChange,
  label,
  block,
  dark,
  value,
  type,
  placeholder,
  disabled
}) {
  return (
    <StyledTextArea
      border={0}
      onChange={onChange}
      p={3}
      m={2}
      borderRadius={1}
      bg={dark ? 'forms.darkBackground' : 'forms.lightBackground'}
      fontSize={1}
      color={dark ? 'forms.lightText' : 'forms.darkText'}
      width={block ? 1 : 2 / 3}
      name={name}
      value={value}
      placeholder={placeholder}
      type={type}
      disabled={disabled}
    >
      {label}
    </StyledTextArea>
  );
}

TextArea.propTypes = {
  name: PropTypes.string,
  onChange: PropTypes.func,
  label: PropTypes.string,
  state: PropTypes.oneOf(['ready', 'valid', 'error']),
  type: PropTypes.string,
  block: PropTypes.bool,
  dark: PropTypes.bool,
  value: PropTypes.string,
  placeholder: PropTypes.string,
  disabled: PropTypes.bool
};
export default TextArea;
