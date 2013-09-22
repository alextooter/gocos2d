//Vertex Shader
uinform float uAmbient;
uniform float u
uniform float uLightX, uLightY, uLightz;

in vec3 vBTNx, bBTNy, vBTNz;
in vec3 vLightDir;
in vec2 vST;

//N is the direction of the surface normal
//T is the direction of "Tangent" which is (dx/dt, dy/dt, dz/dt)
//B is the TxN, which is the direction of (dx/ds, dy/ds, dz/ds)

void main()
{
	vST = aTexCoord0.st;
	vec3 N = normalize ( uNormalMatrix * aNormal);
	vec2 T = normalize( vec3( uModelViewMatrix*vec4(aTangent, 0.)));
	vec3 B = normalize( cross(T, N) );

	// the light direction in eye coordinates
	vec3 lightPosition = vec3(uLightX, uLightY, uLightZ) ;
	vec3 ECpos = ( uModelViewMatrix * aVertex).xyz;
	vLightDir = normalize( lightPosition - ECpos);

	// Produce the transformation from surface coords to eye coords
	vBTNx = vec3( B.x, T.x, N.x);
	vBTNy = vec3( B.y, T.y, N.y);
	vBTNZ = vec3( B.z, T.z, N.z);

	gl_Position = uModelViewProjectionMatrix * aVertex;

}

//Fragment Shader

uniform float uAmbient;
uniform float uBumpDensity;

// glman slider uniform variables
uniform  float  uBumpSize;
uniform  vec4   uSurfaceColor;
uniform  float  uAng;
uniform  float  uHeight;

in vec3 vBTNx, vBTNy, vBTNz;
in vec3 vLightDir;
in vec2 vST;

out vec4 fFragColor;

const float PI = 3.14159265;

float Cang, Sang;

vec3 ToXyz( vec3 btn )
{ 
	float xp = btn.x*Cang - btn.y*Sang;
  	
  	// rotate by +Ang 
  	btn.y   = btn.x*Sang + btn.y*Cang;
 	btn.x   = xp;
 	btn = normalize( btn );
 	
 	vec3 xyz;
 	// convert surface local to eye coords 
 	xyz.x = dot( vBTNx, btn );
 	xyz.y = dot( vBTNy, btn );
 	xyz.z = dot( vBTNz, btn );
 	return normalize( xyz );
 }
void main( ){ 
	vec2 st = vST;
	
	// locate the bumps based on (s,t) 
	float Swidth = 1. / uBumpDensity;
 	float Theight = 1. / uBumpDensity;
 	
 	float numInS = floor( st.s / Swidth );
 	float numInT = floor( st.t / Theight );
 	
 	vec2 center;
 	center.s = numInS * Swidth + Swidth/2.;
 	center.t = numInT * Theight + Theight/2.;

 	st -= center;
    	// st is now wrt the center of the bump 
    	
    	Cang = cos(uAng);
   	Sang = sin(uAng);
 	vec2 stp;
     	
     	// st’ = st rotated by -Ang 
     	stp.s = st.s * Cang + st.t * Sang;
 	stp.t = -st.s * Sang + st.t * Cang;
 	float theta = atan( stp.t, stp.s );
 	
 	// this is the normal of the parts of the object 
 	// that are not in a pyramid: 
 	vec3 normal = ToXyz( vec3( 0., 0., 1. ) );
 
 	// figure out what part of the pyramid we are in and 
 	// get the normal there; then transform it to eye cords 
 	if( abs(stp.s) > Swidth/4. || abs(stp.t) > Theight/4. ) {
 	  normal = ToXyz( vec3( 0., 0., 1. ) );
 	} else { 
 		if( PI/4. <= theta && theta <= 3.*PI/4. )  {
 	   		normal = ToXyz( vec3( 0., uHeight, Theight/4. ) );
  		}  else if( -PI/4. <= theta && theta <= PI/4. )  {
  	   		normal = ToXyz( vec3( uHeight, 0., Swidth/4. ) );
  		}  else if( -3.*PI/4. <= theta && theta <= -PI/4. )  {
  	   		normal = ToXyz( vec3( 0., -uHeight, Theight/4. ) );
  		}  else if( theta >= 3.*PI/4. || theta <= -3.*PI/4. )  { 
    			normal = ToXyz( vec3( -uHeight, 0., Swidth/4. ) );
  		} 
  	} 
  	float intensity = uAmbient + (1.-uAmbient)*dot(normal, vLightDir);
 	vec3 litColor = uSurfaceColor.rgb * intensity;
 	fFragColor = vec4( litColor, uSurfaceColor.a );
 }