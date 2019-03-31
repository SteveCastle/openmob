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

const StyledLabel = styled.label`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function Label({ children, block, dark }) {
  return (
    <StyledLabel
      border={0}
      p={3}
      fontSize={1}
      color={dark ? 'forms.lightText' : 'forms.darkText'}
      width={block ? 1 : 1 / 3}
    >
      {children}
    </StyledLabel>
  );
}

Label.propTypes = {
  children: PropTypes.element,
  state: PropTypes.oneOf(['ready', 'valid', 'error'])
};
export default Label;
