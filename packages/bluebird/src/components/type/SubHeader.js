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

const StyledSubHeader = styled.h2`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function SubHeader({ children, dark }) {
  return (
    <StyledSubHeader
      p={3}
      fontSize={4}
      width={1}
      color={dark ? 'forms.lightText' : 'forms.darkText'}
    >
      {children}
    </StyledSubHeader>
  );
}

SubHeader.propTypes = {
  children: PropTypes.element.isRequired,
  dark: PropTypes.bool
};
export default SubHeader;
