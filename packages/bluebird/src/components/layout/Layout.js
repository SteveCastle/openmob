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

function Layout({ dark, children, tracing }) {
  return (
    <StyledLayout
      display={'flex'}
      border={0}
      py={3}
      flexWrap="wrap"
      width={1}
      bg={
        tracing
          ? `green.${tracing}`
          : dark
          ? 'darkBackground'
          : 'lightBackground'
      }
      fontSize={1}
      color={dark ? 'lightText' : 'darkText'}
    >
      {children}
    </StyledLayout>
  );
}

Layout.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Layout.defaultProps = {};

export default Layout;
