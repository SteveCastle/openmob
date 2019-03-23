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
  `;

function SimpleHero({ title }) {
  return (
    <StyledHero display={'flex'} justifyContent="center" alignItems="center">
      <h1>{title.StringValue}</h1>
    </StyledHero>
  );
}

SimpleHero.propTypes = {
    title: PropTypes.string
  };
  SimpleHero.defaultProps = {};

export default SimpleHero;
