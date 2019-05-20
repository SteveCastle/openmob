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
  display,
  alignItems,
  justifyContent,
  flexWrap,
} from 'styled-system';

const StyledMenuHeader = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${alignItems}
  ${justifyContent}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
`;

function MenuHeader({ children }) {
  return (
    <StyledMenuHeader
      color="gray.0"
      display="flex"
      fontSize={4}
      justifyContent="center"
      alignItems="center"
    >
      {children}
    </StyledMenuHeader>
  );
}

MenuHeader.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
};

export default MenuHeader;
