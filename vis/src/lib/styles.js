/**
 * Styles for what goes inside cytoscape.
 * node and edge weights and styles, etc.
 */

export const colours = {
  // https://materialtheme.arcsine.dev/
  nodeDefault: '#512da8',
  nodeTopicDefault: '#371c8d',
  nodeNearby: '#6242b1',
  edgeDefault: '#797979',
  edgeSelected: '#cbc0e5',
  edgeNearby: '#a29ab7',
  text: '#ffffff',
  bg: '#2c2c2c',
}

export const graphStyles = [
  {
    selector: 'node',
    style: {
      'width': '20',
      'height': '20',
      'font-size': '10',
      'font-weight': 'bold',
      'content': `data(label)`,
      'text-valign': 'center',
      'text-wrap': 'wrap',
      'text-max-width': '140',
      'background-color': colours.nodeDefault,
      'border-color': colours.nodeDefault,
      'border-width': '1',
      'color': colours.text,
      'layoutDimensions': {
        'nodeDimensionsIncludeLabels': true,
      }
    }
  },
  {
    selector: 'node[isTopic]',
    style: {
      'width': '40',
      'height': '40',
      'shape': 'diamond',
      'background-color': colours.nodeDefault,
      'border-color': colours.nodeDefault,
    }
  },
  {
    selector: 'edge',
    style: {
      'curve-style': 'bezier',
      'text-background-color': '#ffffff',
      'text-background-opacity': '1',
      'text-background-padding': '3',
      'width': '3',
      'target-arrow-shape': 'none',
      'line-color': colours.edgeDefault,
      'font-weight': 'bold'
    }
  },
]