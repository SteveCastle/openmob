import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  height,
  fontSize,
  color,
  position,
  borders,
  borderRadius,
  display,
  flexWrap,
  flexDirection,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledContent = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${position}
  ${width}
  ${height}
  ${fontSize}
  ${display}
  ${flexWrap}
  ${flexDirection}
  ${color}
  ${display}
  ${alignItems}
  ${justifyContent}
  overflow-y: auto
`;

function Content({ children, top, left, direction = 'column' }) {
  return (
    <StyledContent
      position="relative"
      display="flex"
      flexDirection={direction}
      alignItems={left ? 'flex-start' : 'center'}
      justifyContent={top ? 'flex-start' : 'center'}
      bg="admin.light.bg"
      border={0}
      flexWrap="wrap"
      width="100%"
      height="100%"
      fontSize={1}
    >
      {children}
    </StyledContent>
  );
}

Content.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  top: PropTypes.bool,
  left: PropTypes.bool,
  direction: PropTypes.string
};
Content.defaultProps = {};

export default Content;
