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

const StyledHeader = styled.h1`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
`;

function Header({ children, dark }) {
  return (
    <StyledHeader
      border={0}
      p={3}
      fontSize={6}
      color={dark ? 'forms.lightText' : 'forms.darkText'}
    >
      {children}
    </StyledHeader>
  );
}

Header.propTypes = {
  children: PropTypes.element.isRequired,
  dark: PropTypes.bool
};
export default Header;
