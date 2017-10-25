/**
 *
 */
package com.jug.openstack.challenges;

import java.io.IOException;
import java.util.Set;

import org.jclouds.ContextBuilder;
import org.jclouds.compute.ComputeService;
import org.jclouds.compute.ComputeServiceContext;
import org.jclouds.compute.domain.Image;
import org.jclouds.logging.slf4j.config.SLF4JLoggingModule;

import com.google.common.collect.ImmutableSet;
import com.google.inject.Module;

/**
 * JCloudImages lists all the images in our <code>openstack-nova</code>
 * @author MarceStarlet
 */
public class JCloudsImages{

	private Set<? extends Image> images;

	public JCloudsImages(){
		Iterable<Module> modules = ImmutableSet.<Module>of(new SLF4JLoggingModule());

		String provider = "openstack-nova";
    String identity = "your-tenantName:your-userName"; // tenantName:userName
    String credential = "your-credential"; // your credential

    ComputeService compute = ContextBuilder.newBuilder(provider)
    		.endpoint("http://8.43.86.2:5000/v2.0")
            .credentials(identity, credential)
            .modules(modules)
            .buildView(ComputeServiceContext.class)
            .getComputeService();

    images = compute.listImages();
	}

  public static void main( String[] args ) throws IOException {
    JCloudsImages jcloudsImages = new JCloudsImages();
    try {
    	jcloudsImages.listImages();
    } catch (Exception e) {
        e.printStackTrace();
    }
  }

  private void listImages() {
    for (Image image : images) {
    	System.out.println("--------------------------------------------------------------------------------------------");
        System.out.println("Image name:        " + image.getName());
        System.out.println("Image description: " + image.getDescription());
        System.out.println("Image Id:          " + image.getId());
        System.out.println("Image OS:          " + image.getOperatingSystem());
        System.out.println("Image location:    " + image.getLocation());
    }
  }
}
