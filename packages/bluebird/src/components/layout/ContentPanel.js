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

const StyledContentPanel = styled.div`
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

function ContentPanel({ children }) {
  return (
    <StyledContentPanel
      display="flex"
      alignItems="center"
      justifyContent="center"
      bg="admin.light.bg"
      border={0}
      flexWrap="wrap"
      width="calc(100% - 256px)"
      height="100%"
      fontSize={1}
    >
      {children}
    </StyledContentPanel>
  );
}

ContentPanel.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ])
};
ContentPanel.defaultProps = {};

export default ContentPanel;
