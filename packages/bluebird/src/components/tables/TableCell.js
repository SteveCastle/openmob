import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  themeGet,
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
  justifyContent,
} from 'styled-system';

const StyledTableCell = styled.td`
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

  &:hover {
      background-color: ${themeGet('colors.cellHover')}
    }

  & a {
    text-decoration: none;
    color: ${themeGet('colors.links')}
  }
 &  a:hover {
      color: ${themeGet('colors.linksHover')}
    }
`;

function TableCell({ children }) {
  return (
    <StyledTableCell px={2} py={3} fontSize={1}>
      {children}
    </StyledTableCell>
  );
}

TableCell.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
};
TableCell.defaultProps = {};

export default TableCell;
