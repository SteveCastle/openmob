import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  maxWidth,
  height,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap,
  flexDirection,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledTableHeader = styled.thead`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${maxWidth}
  ${height}
  ${fontSize}
  ${display}
  ${flexWrap}
  ${flexDirection}
  ${color}
  ${display}
  ${alignItems}
  ${justifyContent}
`;

function TableHeader({ children }) {
  return (
    <StyledTableHeader p={2} border={0} fontSize={1}>
      {children}
    </StyledTableHeader>
  );
}

TableHeader.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element)
};
TableHeader.defaultProps = {};

export default TableHeader;
