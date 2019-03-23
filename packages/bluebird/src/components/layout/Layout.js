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
  flexWrap
} from 'styled-system';

const StyledLayout = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${flexWrap}
  ${color}
  ${display}
`;

function Layout({ children }) {
  return (
    <StyledLayout
      display={'flex'}
      border={0}
      flexWrap="wrap"
      width={1}
      fontSize={1}
    >
      {children}
    </StyledLayout>
  );
}

Layout.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
};
Layout.defaultProps = {};

export default Layout;
