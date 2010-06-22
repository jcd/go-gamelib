package main

import "gl"
import "game/phys"
import . "game/common"

var rtri gl.GLfloat = 0.0;
var rquad gl.GLfloat = 0.0;

func drawCube(v *Vector3) { 

	gl.LoadIdentity()
	//	gl.Translatef( -1.5, 0.0, -6.0 )
	gl.Translatef( gl.GLfloat(v.X), gl.GLfloat(v.Y), gl.GLfloat(v.Z), )
	
	/* Rotate The Triangle On The Y axis ( NEW ) */
	// gl.Rotatef( rtri, 0.0, 1.0, 0.0 );

	gl.Begin( gl.TRIANGLES );             /* Drawing Using Triangles       */
	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Front)       */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f( -1.0, -1.0,  1.0 ); /* Left Of Triangle (Front)      */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f(  1.0, -1.0,  1.0 ); /* Right Of Triangle (Front)     */

	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Right)       */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f(  1.0, -1.0,  1.0 ); /* Left Of Triangle (Right)      */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f(  1.0, -1.0, -1.0 ); /* Right Of Triangle (Right)     */

	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Back)        */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f(  1.0, -1.0, -1.0 ); /* Left Of Triangle (Back)       */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f( -1.0, -1.0, -1.0 ); /* Right Of Triangle (Back)      */

	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Left)        */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f( -1.0, -1.0, -1.0 ); /* Left Of Triangle (Left)       */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f( -1.0, -1.0,  1.0 ); /* Right Of Triangle (Left)      */
	gl.End( );                            /* Finished Drawing The Triangle */

	rtri = rtri + 0.2;

}

func drawMass(m *phys.Mass, dsz gl.GLfloat) { 

	gl.LoadIdentity()
	//	gl.Translatef( -1.5, 0.0, -6.0 )
	
	v := m.Pos
	gl.Translatef( gl.GLfloat(v.X), gl.GLfloat(v.Y), gl.GLfloat(v.Z), )
	
	/* Rotate The Triangle On The Y axis ( NEW ) */
	// gl.Rotatef( rtri, 0.0, 1.0, 0.0 );

	gl.Begin( gl.TRIANGLES );             /* Drawing Using Triangles       */
	gl.Color3f(   dsz,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  dsz,  0.0 ); /* Top Of Triangle (Front)       */
	gl.Color3f(   0.0,  dsz,  0.0 ); /* Green                         */
	gl.Vertex3f( -dsz, -dsz,  dsz ); /* Left Of Triangle (Front)      */
	gl.Color3f(   0.0,  0.0,  dsz ); /* Blue                          */
	gl.Vertex3f(  dsz, -dsz,  dsz ); /* Right Of Triangle (Front)     */

	gl.Color3f(   dsz,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  dsz,  0.0 ); /* Top Of Triangle (Right)       */
	gl.Color3f(   0.0,  0.0,  dsz ); /* Blue                          */
	gl.Vertex3f(  dsz, -dsz,  dsz ); /* Left Of Triangle (Right)      */
	gl.Color3f(   0.0,  dsz,  0.0 ); /* Green                         */
	gl.Vertex3f(  dsz, -dsz, -dsz ); /* Right Of Triangle (Right)     */

	gl.Color3f(   dsz,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  dsz,  0.0 ); /* Top Of Triangle (Back)        */
	gl.Color3f(   0.0,  dsz,  0.0 ); /* Green                         */
	gl.Vertex3f(  dsz, -dsz, -dsz ); /* Left Of Triangle (Back)       */
	gl.Color3f(   0.0,  0.0,  dsz ); /* Blue                          */
	gl.Vertex3f( -dsz, -dsz, -dsz ); /* Right Of Triangle (Back)      */

	gl.Color3f(   dsz,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  dsz,  0.0 ); /* Top Of Triangle (Left)        */
	gl.Color3f(   0.0,  0.0,  dsz ); /* Blue                          */
	gl.Vertex3f( -dsz, -dsz, -dsz ); /* Left Of Triangle (Left)       */
	gl.Color3f(   0.0,  dsz,  0.0 ); /* Green                         */
	gl.Vertex3f( -dsz, -dsz,  dsz ); /* Right Of Triangle (Left)      */
	gl.End( );                            /* Finished Drawing The Triangle */

}

func drawCubes(v *Vector3) { 

	gl.LoadIdentity()
	//	gl.Translatef( -1.5, 0.0, -6.0 )
	gl.Translatef( gl.GLfloat(v.X), gl.GLfloat(v.Y), gl.GLfloat(v.Z), )
	
	/* Rotate The Triangle On The Y axis ( NEW ) */
	gl.Rotatef( rtri, 0.0, 1.0, 0.0 );

	gl.Begin( gl.TRIANGLES );             /* Drawing Using Triangles       */
	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Front)       */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f( -1.0, -1.0,  1.0 ); /* Left Of Triangle (Front)      */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f(  1.0, -1.0,  1.0 ); /* Right Of Triangle (Front)     */

	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Right)       */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f(  1.0, -1.0,  1.0 ); /* Left Of Triangle (Right)      */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f(  1.0, -1.0, -1.0 ); /* Right Of Triangle (Right)     */

	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Back)        */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f(  1.0, -1.0, -1.0 ); /* Left Of Triangle (Back)       */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f( -1.0, -1.0, -1.0 ); /* Right Of Triangle (Back)      */

	gl.Color3f(   1.0,  0.0,  0.0 ); /* Red                           */
	gl.Vertex3f(  0.0,  1.0,  0.0 ); /* Top Of Triangle (Left)        */
	gl.Color3f(   0.0,  0.0,  1.0 ); /* Blue                          */
	gl.Vertex3f( -1.0, -1.0, -1.0 ); /* Left Of Triangle (Left)       */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Green                         */
	gl.Vertex3f( -1.0, -1.0,  1.0 ); /* Right Of Triangle (Left)      */
	gl.End( );                            /* Finished Drawing The Triangle */

	rtri = rtri + 0.2;

	return

	/* Move Right 3 Units */
	gl.LoadIdentity( );
	gl.Translatef( 1.5, 0.0, -6.0 );

	/* Rotate The Quad On The X axis ( NEW ) */
	gl.Rotatef( rquad, 1.0, 0.0, 0.0 );

	/* Set The Color To Blue One Time Only */
	gl.Color3f( 0.5, 0.5, 1.0);

	gl.Begin( gl.QUADS );                 /* Draw A Quad                      */
	gl.Color3f(   0.0,  1.0,  0.0 ); /* Set The Color To Green           */
	gl.Vertex3f(  1.0,  1.0, -1.0 ); /* Top Right Of The Quad (Top)      */
	gl.Vertex3f( -1.0,  1.0, -1.0 ); /* Top Left Of The Quad (Top)       */
	gl.Vertex3f( -1.0,  1.0,  1.0 ); /* Bottom Left Of The Quad (Top)    */
	gl.Vertex3f(  1.0,  1.0,  1.0 ); /* Bottom Right Of The Quad (Top)   */

	gl.Color3f(   1.0,  0.5,  0.0 ); /* Set The Color To Orange          */
	gl.Vertex3f(  1.0, -1.0,  1.0 ); /* Top Right Of The Quad (Botm)     */
	gl.Vertex3f( -1.0, -1.0,  1.0 ); /* Top Left Of The Quad (Botm)      */
	gl.Vertex3f( -1.0, -1.0, -1.0 ); /* Bottom Left Of The Quad (Botm)   */
	gl.Vertex3f(  1.0, -1.0, -1.0 ); /* Bottom Right Of The Quad (Botm)  */

	gl.Color3f(   1.0,  0.0,  0.0 ); /* Set The Color To Red             */
	gl.Vertex3f(  1.0,  1.0,  1.0 ); /* Top Right Of The Quad (Front)    */
	gl.Vertex3f( -1.0,  1.0,  1.0 ); /* Top Left Of The Quad (Front)     */
	gl.Vertex3f( -1.0, -1.0,  1.0 ); /* Bottom Left Of The Quad (Front)  */
	gl.Vertex3f(  1.0, -1.0,  1.0 ); /* Bottom Right Of The Quad (Front) */

	gl.Color3f(   1.0,  1.0,  0.0 ); /* Set The Color To Yellow          */
	gl.Vertex3f(  1.0, -1.0, -1.0 ); /* Bottom Left Of The Quad (Back)   */
	gl.Vertex3f( -1.0, -1.0, -1.0 ); /* Bottom Right Of The Quad (Back)  */
	gl.Vertex3f( -1.0,  1.0, -1.0 ); /* Top Right Of The Quad (Back)     */
	gl.Vertex3f(  1.0,  1.0, -1.0 ); /* Top Left Of The Quad (Back)      */

	gl.Color3f(   0.0,  0.0,  1.0 ); /* Set The Color To Blue            */
	gl.Vertex3f( -1.0,  1.0,  1.0 ); /* Top Right Of The Quad (Left)     */
	gl.Vertex3f( -1.0,  1.0, -1.0 ); /* Top Left Of The Quad (Left)      */
	gl.Vertex3f( -1.0, -1.0, -1.0 ); /* Bottom Left Of The Quad (Left)   */
	gl.Vertex3f( -1.0, -1.0,  1.0 ); /* Bottom Right Of The Quad (Left)  */

	gl.Color3f(   1.0,  0.0,  1.0 ); /* Set The Color To Violet          */
	gl.Vertex3f(  1.0,  1.0, -1.0 ); /* Top Right Of The Quad (Right)    */
	gl.Vertex3f(  1.0,  1.0,  1.0 ); /* Top Left Of The Quad (Right)     */
	gl.Vertex3f(  1.0, -1.0,  1.0 ); /* Bottom Left Of The Quad (Right)  */
	gl.Vertex3f(  1.0, -1.0, -1.0 ); /* Bottom Right Of The Quad (Right) */
	gl.End( );                            /* Done Drawing The Quad            */

	rquad -= 0.15;
}
