import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  fontFamily,
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
  ${fontFamily}
  ${color}
`;

function Header({ children, dark }) {
  return (
    <StyledHeader
      p={3}
      pb={0}
      mb={0}
      fontSize={6}
      fontFamily='sansSerif'
      width={1}
      color={dark ? 'violet.0' : 'gray.9'}
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
