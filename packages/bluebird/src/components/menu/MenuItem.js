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
  flexWrap
} from 'styled-system';

const StyledItem = styled.div`
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

function MenuItem({ title, icon, action }) {
  return (
    <StyledItem
      color="gray.0"
      display="flex"
      p={2}
      fontSize={2}
      justifyContent="center"
      alignItems="center"
    >
      {title}
    </StyledItem>
  );
}

MenuItem.propTypes = {
  title: PropTypes.string,
  icon: PropTypes.string,
  action: PropTypes.func
};

export default MenuItem;
