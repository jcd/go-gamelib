package main

import "gl"
import "sdl/ttf"
// import "math"
import "sdl"
import . "game/common"

import "unsafe"
//import "fmt"

type Font struct {
	Height float	
	font *ttf.Font
//	Textures *gl.uint // texture ids
//	Listbase gl.uint // id of the first display list id
}

func NewFont(fname string, height int) *Font {
//	f.Textures = [12]gl.uint
	
	f := &Font{float(height), ttf.OpenFont(fname, height)}

	if f.font == nil {
		panic("Couldn't open font " + fname)
	}
	return f
}

func next_p2(a int) int {
	var rval int = 1;
	for rval<a { rval<<=1 }
	return rval;
}

func (f *Font) DrawText2D(x int, y int, color [3]byte, scale float, txt string) (endw int, endh int) {
	gl.LoadIdentity();
//	gl.Translatef(0.0, 0.0, -2.0);

	glEnable2D()
	gl.Disable(gl.DEPTH_TEST)

	texture, initial, intermediarya, intermediary, w, h  := f.setupTextRendering(color, txt)

	locX := gl.GLfloat(x)
	locY := gl.GLfloat(y)
	wi := gl.GLfloat(w) * gl.GLfloat(scale)
	he := gl.GLfloat(h) * gl.GLfloat(scale)

	/* Draw a quad at location */
	gl.Begin(gl.QUADS);
		/* Recall that the origin is in the lower-left corner
		   That is why the TexCoords specify different corners
		   than the Vertex coors seem to. */
		gl.TexCoord2f(0.0, 1.0); 
			gl.Vertex2f(locX    , locY)
		gl.TexCoord2f(1.0, 1.0); 
			gl.Vertex2f(locX + wi, locY)
		gl.TexCoord2f(1.0, 0.0); 
			gl.Vertex2f(locX + wi, locY + he)
		gl.TexCoord2f(0.0, 0.0); 
			gl.Vertex2f(locX     , locY + he)
	gl.End();

	endw, endh = f.teardownTextRendering(texture, initial, intermediarya, intermediary)

	gl.Enable(gl.DEPTH_TEST)
	glDisable2D()
	return
}

func (f *Font) DrawText3D(pos Vector3, color [3]byte, scale float, txt string) (endw int, endh int) {

	gl.LoadIdentity();
// 	gl.Translatef(gl.GLfloat(pos.X), gl.GLfloat(pos.Y), gl.GLfloat(-4.99));
 	gl.Translatef(gl.GLfloat(pos.X), gl.GLfloat(pos.Y), gl.GLfloat(-4.99));
 	gl.Scalef(gl.GLfloat(1.0/64.0 * scale), gl.GLfloat(1.0/64.0 * scale), gl.GLfloat(1.0));
	gl.Disable(gl.DEPTH_TEST)

	texture, initial, intermediarya, intermediary, w ,h  := f.setupTextRendering(color, txt)

	wi := gl.GLfloat(w) 
	he := gl.GLfloat(h) 
	locX := gl.GLfloat(0) - wi * 0.5
	locY := gl.GLfloat(0) - he * 0.5

	/* Draw a quad at location */
	gl.Begin(gl.QUADS);
		/* Recall that the origin is in the lower-left corner
		   That is why the TexCoords specify different corners
		   than the Vertex coors seem to. */

	        gl.Color4f(1.0, 1.0, 0.0, 1.0)

		gl.TexCoord2f(0.0, 1.0); 
			gl.Vertex3f(locX     , locY, 0.0)
		gl.TexCoord2f(1.0, 1.0); 
			gl.Vertex3f(locX + wi , locY, 0.0)
		gl.TexCoord2f(1.0, 0.0); 
			gl.Vertex3f(locX + wi, locY + he, 0.0)
		gl.TexCoord2f(0.0, 0.0); 
			gl.Vertex3f(locX     , locY + he, 0.0)
	gl.End();

	endw, endh = f.teardownTextRendering(texture, initial, intermediarya, intermediary)

	gl.Enable(gl.DEPTH_TEST)
	return
}

func (f *Font) setupTextRendering(color [3]byte, txt string) (gl.GLuint, *sdl.Surface, *sdl.Surface, *sdl.Surface, int, int) {

	//
	var texture gl.GLuint
	
	/* Use SDL_TTF to render our text */
	var col sdl.Color
	col.R = color[0]
	col.G = color[1]
	col.B = color[2]

	// get surface with text
	initial := ttf.RenderText_Blended(f.font, txt, col)
	
	/* Convert the rendered text to a known format */
	w := next_p2(int(initial.W))
	h := next_p2(int(initial.H))
	
	intermediarya := sdl.CreateRGBSurface(0, w, h, 32, 
		0x00ff0000, 0x0000ff00, 0x000000ff, 0xff000000);

	var rr *sdl.Rect = nil
	intermediarya.Blit(rr, initial, rr);

	intermediary := intermediarya.DisplayFormatAlpha()
	
	/* Tell GL about our new texture */
	gl.GenTextures(1, &texture);
	gl.BindTexture(gl.TEXTURE_2D, texture);
	gl.TexImage2D(gl.TEXTURE_2D, 0, 4, gl.GLsizei(w), gl.GLsizei(h), 0, gl.BGRA, 
		gl.UNSIGNED_BYTE, unsafe.Pointer(intermediary.Pixels) );
	

	gl.TexEnvi( gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE );

	/* GL_NEAREST looks horrible, if scaled... */
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);	

	/* prepare to render our texture */
	gl.Enable(gl.TEXTURE_2D);

	gl.BindTexture(gl.TEXTURE_2D, texture);
	gl.Color4f(1.0, 1.0, 1.0, 1.0);

	gl.Enable(gl.BLEND)
//	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.BlendFunc(gl.ONE, gl.ONE_MINUS_SRC_COLOR) // TODO : MAke real alpha work!

	return texture, initial, intermediarya, intermediary, w, h
}

func (f *Font) teardownTextRendering(texture gl.GLuint, initial *sdl.Surface, intermediarya *sdl.Surface, intermediary *sdl.Surface) (endw int, endh int) {

	// 
	gl.Disable(gl.BLEND)

	/* Bad things happen if we delete the texture before it finishes */
	gl.Finish();
	
	/* return the deltas in the unused w,h part of the rect */
	endw = int(initial.W)
	endh = int(initial.H)
	
	/* Clean up */
	initial.Free()
	intermediarya.Free()
	intermediary.Free()
	gl.DeleteTextures(1, &texture);
	//
	return
}

func glEnable2D() {

	var vPort [4]gl.GLint
  
	gl.GetIntegerv(gl.VIEWPORT, &vPort[0]);
  
	gl.MatrixMode(gl.PROJECTION);
	gl.PushMatrix();
	gl.LoadIdentity();
  
	gl.Ortho(0.0, gl.GLdouble(vPort[2]), 0.0, gl.GLdouble(vPort[3]), -1.0, 1.0);
	gl.MatrixMode(gl.MODELVIEW);
	gl.PushMatrix();
	gl.LoadIdentity();
}

func glDisable2D() {
	gl.MatrixMode(gl.PROJECTION);
	gl.PopMatrix();   
	gl.MatrixMode(gl.MODELVIEW);
	gl.PopMatrix();	
}
