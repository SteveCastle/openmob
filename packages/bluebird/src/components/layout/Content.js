import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  height,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap,
  flexDirection,
  alignItems,
  justifyContent,
} from 'styled-system';

const StyledContent = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
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

function Content({ children, top, left }) {
  return (
    <StyledContent
      display="flex"
      flexDirection="column"
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
    PropTypes.node,
  ]),
  top: PropTypes.bool,
  left: PropTypes.bool,
};
Content.defaultProps = {};

export default Content;
