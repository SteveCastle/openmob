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
  borderBottom,
  borderRadius,
  display,
  flexWrap,
  flexDirection,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledTableRow = styled.tr`
  ${borderBottom}
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

function TableRow({ children }) {
  return (
    <StyledTableRow p={2} fontSize={1}>
      {children}
    </StyledTableRow>
  );
}

TableRow.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element)
};
TableRow.defaultProps = {};

export default TableRow;
