import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  background,
  fontSize,
  color,
  borders,
  position,
  bottom,
  borderRadius
} from 'styled-system';
import Card from '../cards/Card';

const StyledItem = styled.li`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${position}
  ${fontSize}
  ${color}
  box-sizing: border-box;
  cursor: pointer;
  transition: 2s;
  :hover {
    opacity: .7
  }
`;

const StyledTitle = styled.h2`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
`;

const CardBody = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
`;

const CardLabel = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
  ${background}
  height: 200px;
  background-image: url("${props => props.image}");
  background-position: center;
  background-origin: content-box;
  background-repeat: no-repeat;
  background-size: cover;
`;

const StyledCaption = styled.p`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
`;

function GridItem({ width, title, caption, onClick, image }) {
  return (
    <StyledItem
      p={2}
      m={2}
      position="relative"
      onClick={onClick}
      width={[1, 1, width / 12]}
    >
      <Card fillHeight>
        <CardLabel width={1} image={image} />
        <CardBody>
          <StyledTitle color="gray.9" mx={3} my={3}>
            {title}
          </StyledTitle>
          <StyledCaption color="gray.9" mx={3} my={2}>
            {caption}
          </StyledCaption>
        </CardBody>
      </Card>
    </StyledItem>
  );
}

GridItem.propTypes = {
  width: PropTypes.number,
  title: PropTypes.string,
  caption: PropTypes.string,
  image: PropTypes.string,
  onClick: PropTypes.func
};
GridItem.defaultProps = { width: 3 };

export default GridItem;
