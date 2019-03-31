import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borderBottom,
  borderRadius,
  display,
  alignItems,
  justifyContent,
  flexWrap
} from 'styled-system';

const StyledItem = styled.div`
  ${borderBottom}
  ${borderRadius}
  ${space}
  ${width}
  ${alignItems}
  ${justifyContent}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  a {
    text-decoration: none;
    color: white;
    height: 100%;
    width: 100%;
    padding: 16px 0 16px 16px;
  }
`;

function MenuItem({ children }) {
  return (
    <StyledItem
      borderBottom="1px solid #3f3f73"
      color="gray.0"
      display="flex"
      fontSize={2}
      justifyContent="flex-start"
      alignItems="center"
    >
      {children}
    </StyledItem>
  );
}

MenuItem.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ])
};

export default MenuItem;
