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
  alignItems,
  justifyContent
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
  ${color}
  ${display}
  ${alignItems}
  ${justifyContent}
`;

function Content({ children }) {
  return (
    <StyledContent
      display="flex"
      alignItems="center"
      justifyContent="center"
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
  children: PropTypes.arrayOf(PropTypes.element)
};
Content.defaultProps = {};

export default Content;
