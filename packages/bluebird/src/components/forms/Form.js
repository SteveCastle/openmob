import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
} from 'styled-system';

const StyledForm = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function Form({ children, onChange, state }) {
  return (
    <StyledForm border={0} p={3} fontSize={1}>
      {children}
    </StyledForm>
  );
}

Form.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
  onChange: PropTypes.func,
  state: PropTypes.oneOf(['ready', 'valid', 'error']),
};
export default Form;
