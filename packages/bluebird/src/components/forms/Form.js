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

const StyledForm = styled.form`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function Form({ onChange, state }) {
  return <StyledForm border={0} p={3} fontSize={1} />;
}

Form.propTypes = {
  onChange: PropTypes.func,
  state: PropTypes.oneOf(['ready', 'valid', 'error'])
};
export default Form;
