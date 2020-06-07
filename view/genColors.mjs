import rcolor from 'rcolor'
import fs from "fs"


const w = fs.createWriteStream( "./colors.json" );

const colors = new Array( 300 ).fill().map( v => rcolor() );

w.write( JSON.stringify(
  Array.from( new Set( colors ) ),
  null,
  4
) );

