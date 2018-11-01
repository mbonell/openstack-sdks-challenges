/**
 *
 */
package com.jug.openstack.challenges;

import java.io.IOException;

import org.jclouds.ContextBuilder;
import org.jclouds.compute.ComputeService;
import org.jclouds.compute.ComputeServiceContext;
import org.jclouds.compute.domain.Hardware;
import org.jclouds.compute.domain.NodeMetadata;
import org.jclouds.logging.slf4j.config.SLF4JLoggingModule;

import com.google.common.collect.ImmutableSet;
import com.google.inject.Module;

/**
 * JCloudsFlavors lists hardware flavors from our <code>openstack-nova</code>
 * @author MarceStarlet
 */
public class JCloudsFlavors {

	private NodeMetadata metadata;

	public JCloudsFlavors(){
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

    metadata = compute.getNodeMetadata("RegionOne/e2d1ccd8-a939-4d72-b4c0-a8699a7c9286"); // use your region metadata
    System.out.println(metadata);
	}

	public static void main( String[] args ) throws IOException {
    JCloudsFlavors jcloudsFlavors = new JCloudsFlavors();
    try {
    	jcloudsFlavors.listHardware();
    } catch (Exception e) {
        e.printStackTrace();
    }
  }

	private void listHardware() {
		Hardware hw = metadata.getHardware();
  	System.out.println("--------------------------------------------------------------------------------------------");
    System.out.println("Hardware name:        " + hw.getName());
    System.out.println("Hardware id:          " + hw.getId());
    System.out.println("Hardware id:          " + hw.getProcessors());
    System.out.println("Hardware ram:         " + hw.getRam());
    System.out.println("Hardware type:        " + hw.getType());
    System.out.println("Hardware location:    " + hw.getLocation());
  }

}
