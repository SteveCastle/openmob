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
  borderRadius,
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
      m={0}
      fontSize={6}
      fontFamily="sansSerif"
      width={1}
      color={dark ? 'type.dark.header' : 'type.light.header'}
    >
      {children}
    </StyledHeader>
  );
}

Header.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]).isRequired,
  dark: PropTypes.bool,
};
export default Header;
