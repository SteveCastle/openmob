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
const Container = styled.div`
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
const StyledDataTable = styled.table`
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
  table-layout:fixed;
  width:100%;

  & td {
    word-wrap: break-word;
  }
`;

function DataTable({ children }) {
  return (
    <Container>
      <StyledDataTable p={2} border={0} fontSize={1}>
        {children}
      </StyledDataTable>
    </Container>
  );
}

DataTable.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ])
};
DataTable.defaultProps = {};

export default DataTable;
