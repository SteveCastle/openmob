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

const StyledSignups = styled.section`
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

function SimpleSignups({ title }) {
  return (
    <StyledSignups
      display={'flex'}
      justifyContent="center"
      alignItems="center"
      bg="fuschia.8"
    >
      <Title color="gray.0" size={1}>
        {title}
      </Title>
    </StyledSignups>
  );
}

SimpleSignups.propTypes = {
  title: PropTypes.string.isRequired
};
SimpleSignups.defaultProps = { title: 'Simple Signups' };

export default SimpleSignups;
