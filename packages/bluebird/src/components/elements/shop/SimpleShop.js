import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';

import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledShop = styled.section`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  min-height: 400px;
  `;
const Title = styled.h1`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${alignItems}
  ${justifyContent}
  ${color}
  ${display}
  `;

function SimpleShop({ title }) {
  return (
    <StyledShop
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="violet.7"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledShop>
  );
}

SimpleShop.propTypes = {
  title: PropTypes.string.isRequired
};
SimpleShop.defaultProps = { title: 'Simple Shop' };

export default SimpleShop;
