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

const StyledHero = styled.section`
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
  background: #3494E6;  /* fallback for old browsers */
  background: -webkit-linear-gradient(to right, #EC6EAD, #3494E6);
  background: linear-gradient(to right, #EC6EAD, #3494E6); 
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

function SimpleHero({ title }) {
  return (
    <StyledHero display={'flex'} justifyContent="center" alignItems="center">
      <Title color="gray.0" size={4}>
        {title}
      </Title>
    </StyledHero>
  );
}

SimpleHero.propTypes = {
  title: PropTypes.string.isRequired
};
SimpleHero.defaultProps = {};

export default SimpleHero;
